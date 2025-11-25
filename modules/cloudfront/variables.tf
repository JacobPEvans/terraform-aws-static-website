variable "domain_name" {
  description = "Main domain name"
  type        = string
}

variable "aliases" {
  description = "Aliases for the CloudFront distribution"
  type        = list(string)
}

variable "price_class" {
  description = "CloudFront price class"
  type        = string
  default     = "PriceClass_All"
}

variable "root_bucket_id" {
  description = "ID of the root S3 bucket"
  type        = string
}

variable "root_bucket_website_endpoint" {
  description = "Website endpoint of the root S3 bucket"
  type        = string
}

variable "redirect_bucket_id" {
  description = "ID of the redirect S3 bucket"
  type        = string
}

variable "redirect_bucket_website_endpoint" {
  description = "Website endpoint of the redirect S3 bucket"
  type        = string
}

variable "logs_bucket_domain_name" {
  description = "Domain name of the logs S3 bucket"
  type        = string
}

variable "acm_certificate_arn" {
  description = "ARN of the ACM certificate"
  type        = string
}

variable "support_spa" {
  description = "Whether to support Single Page Application routing"
  type        = bool
  default     = false
}

variable "lambda_function_arn" {
  description = "ARN of Lambda@Edge function"
  type        = string
  default     = null
}

variable "lambda_function_event_type" {
  description = "Event type for Lambda@Edge function"
  type        = string
  default     = "origin-request"
}

variable "redirect_enabled" {
  description = "Whether redirect distribution is enabled"
  type        = bool
  default     = true
}

variable "redirect_domain" {
  description = "Domain for redirect distribution"
  type        = string
  default     = ""
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {}
}
