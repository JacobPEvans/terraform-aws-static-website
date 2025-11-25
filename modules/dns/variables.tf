variable "route53_zone_id" {
  description = "Route53 hosted zone ID"
  type        = string
}

variable "domain_name" {
  description = "Main domain name"
  type        = string
}

variable "redirect_domain" {
  description = "Redirect domain name"
  type        = string
  default     = ""
}

variable "redirect_enabled" {
  description = "Whether redirect is enabled"
  type        = bool
  default     = true
}

variable "root_distribution_domain_name" {
  description = "Domain name of the root CloudFront distribution"
  type        = string
}

variable "root_distribution_hosted_zone_id" {
  description = "Hosted zone ID of the root CloudFront distribution"
  type        = string
}

variable "redirect_distribution_domain_name" {
  description = "Domain name of the redirect CloudFront distribution"
  type        = string
  default     = ""
}

variable "redirect_distribution_hosted_zone_id" {
  description = "Hosted zone ID of the redirect CloudFront distribution"
  type        = string
  default     = ""
}
