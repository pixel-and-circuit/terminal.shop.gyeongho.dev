variable "aws_region" {
  description = "AWS region for Lightsail instance"
  type        = string
  default     = "ap-northeast-2"
}

variable "instance_name" {
  description = "Name of the Lightsail instance"
  type        = string
  default     = "shop-terminal"
}

variable "admin_ssh_allowed_cidrs" {
  description = "CIDR blocks allowed to access management SSH (port 2222). Find your IP: curl -s https://checkip.amazonaws.com"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "key_pair_name" {
  description = "Lightsail key pair name for SSH access. Create with: aws lightsail import-key-pair"
  type        = string
  default     = ""
}
