variable "website-domain-main" {
  description = "Main website domain (e.g., example.com)"
  type        = string
}

variable "website-domain-redirect" {
  description = "Domain to redirect to main domain (e.g., www.example.com)"
  type        = string
  default     = ""
}

variable "domains-zone-root" {
  description = "Root domain for Route53 hosted zone"
  type        = string
}

variable "support-spa" {
  description = "Enable SPA support (404 redirects to index.html)"
  type        = bool
  default     = false
}
