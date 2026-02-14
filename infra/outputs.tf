output "instance_public_ip" {
  description = "Static IP — Cloudflare DNS A record에 등록할 주소"
  value       = aws_lightsail_static_ip.shop.ip_address
}

output "ssh_admin_command" {
  description = "관리용 SSH 접속 명령어"
  value       = "ssh -p 2222 ubuntu@${aws_lightsail_static_ip.shop.ip_address}"
}
