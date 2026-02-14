#!/bin/bash
set -euxo pipefail

export DEBIAN_FRONTEND=noninteractive
apt-get update && apt-get upgrade -y
apt-get install -y ufw fail2ban

# sshd를 2222 포트로 이동 (22는 shop이 사용)
sed -i 's/^#\?Port 22$/Port 2222/' /etc/ssh/sshd_config
grep -q '^Port 2222' /etc/ssh/sshd_config || echo "Port 2222" >> /etc/ssh/sshd_config
systemctl restart sshd

# 방화벽
ufw default deny incoming
ufw default allow outgoing
ufw allow 22/tcp comment 'Shop SSH TUI'
ufw allow 2222/tcp comment 'Management SSH'
ufw --force enable

# shop 전용 유저 및 디렉토리
useradd -r -s /bin/false -d /opt/shop shop || true
mkdir -p /opt/shop/bin /opt/shop/.ssh
chown -R shop:shop /opt/shop

# systemd 서비스
cat > /etc/systemd/system/shop.service <<'EOF'
[Unit]
Description=Terminal Shop SSH Server
After=network.target

[Service]
Type=simple
User=shop
Group=shop
WorkingDirectory=/opt/shop
ExecStart=/opt/shop/bin/shop --ssh
Restart=always
RestartSec=5
Environment="SHOP_SSH=1"
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/shop
MemoryMax=400M
TasksMax=100

# 포트 22 바인딩 권한
AmbientCapabilities=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload

# fail2ban (관리용 SSH 보호)
cat > /etc/fail2ban/jail.local <<'EOF'
[sshd]
enabled = true
port = 2222
maxretry = 5
bantime = 3600
EOF
systemctl enable fail2ban && systemctl restart fail2ban

echo "user_data complete"
