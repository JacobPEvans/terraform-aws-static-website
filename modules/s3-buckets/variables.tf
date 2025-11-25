variable "domain_name" {
  description = "Main domain name for the website"
  type        = string
}

variable "support_spa" {
  description = "Whether to support Single Page Application routing"
  type        = bool
  default     = false
}

variable "force_destroy" {
  description = "Allow destruction of non-empty buckets"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default     = {}
}
