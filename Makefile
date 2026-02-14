.PHONY: format build test run pre-commit-install
.PHONY: build-linux deploy deploy-binary deploy-restart deploy-status deploy-logs deploy-ssh
.PHONY: infra-init infra-plan infra-apply infra-destroy infra-output

format:
	gofmt -s -w .
	@which goimports >/dev/null 2>&1 && goimports -w . || true

build:
	go build -o bin/shop ./cmd/shop

test:
	go test ./...

run:
	go run ./cmd/shop

pre-commit-install:
	@which pre-commit >/dev/null 2>&1 && pre-commit install || echo "pre-commit not installed; run: pip install pre-commit"

# ---------------------------------------------------------------------------
# Cross-compile
# ---------------------------------------------------------------------------

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/shop-linux ./cmd/shop

# ---------------------------------------------------------------------------
# Infrastructure (Terraform)
# ---------------------------------------------------------------------------

infra-init:
	cd infra && terraform init

infra-plan:
	cd infra && terraform plan

infra-apply:
	cd infra && terraform apply

infra-destroy:
	@echo "WARNING: This will destroy all infrastructure!"
	@read -p "Type 'yes' to confirm: " c && [ "$$c" = "yes" ]
	cd infra && terraform destroy

infra-output:
	cd infra && terraform output

# ---------------------------------------------------------------------------
# Deployment
# ---------------------------------------------------------------------------

LIGHTSAIL_IP ?= $(shell cd infra && terraform output -raw instance_public_ip 2>/dev/null)
SSH_CMD = ssh -p 2222 -o StrictHostKeyChecking=no ubuntu@$(LIGHTSAIL_IP)

deploy: build-linux deploy-binary deploy-restart
	@echo "Done. Test: ssh shop.gyeongho.dev"

deploy-binary:
	scp -P 2222 -o StrictHostKeyChecking=no bin/shop-linux ubuntu@$(LIGHTSAIL_IP):/tmp/shop
	$(SSH_CMD) 'sudo mv /tmp/shop /opt/shop/bin/shop && sudo chown shop:shop /opt/shop/bin/shop && sudo chmod +x /opt/shop/bin/shop'

deploy-restart:
	$(SSH_CMD) 'sudo systemctl enable shop && sudo systemctl restart shop'
	@sleep 2
	@$(MAKE) deploy-status

deploy-status:
	$(SSH_CMD) 'sudo systemctl status shop --no-pager'

deploy-logs:
	$(SSH_CMD) 'sudo journalctl -u shop -n 50 --no-pager'

deploy-ssh:
	ssh -p 2222 ubuntu@$(LIGHTSAIL_IP)
