terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_lightsail_instance" "shop" {
  name              = var.instance_name
  availability_zone = "${var.aws_region}a"
  blueprint_id      = "ubuntu_22_04"
  bundle_id         = "nano_3_0"
  key_pair_name     = var.key_pair_name != "" ? var.key_pair_name : null

  user_data = file("${path.module}/user_data.sh")

  tags = {
    Name      = "shop-terminal"
    ManagedBy = "terraform"
  }
}

resource "aws_lightsail_static_ip" "shop" {
  name = "${var.instance_name}-ip"
}

resource "aws_lightsail_static_ip_attachment" "shop" {
  static_ip_name = aws_lightsail_static_ip.shop.name
  instance_name  = aws_lightsail_instance.shop.name
}

resource "aws_lightsail_instance_public_ports" "shop" {
  instance_name = aws_lightsail_instance.shop.name

  port_info {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
    cidrs     = ["0.0.0.0/0"]
  }

  port_info {
    protocol  = "tcp"
    from_port = 2222
    to_port   = 2222
    cidrs     = var.admin_ssh_allowed_cidrs
  }
}
