variable "domain_root" {
  description = "Root domain for the certificate"
  type        = string
}

variable "route53_zone_id" {
  description = "Route53 hosted zone ID for DNS validation"
  type        = string
}

variable "tags" {
  description = "Tags to apply to the certificate"
  type        = map(string)
  default     = {}
}
