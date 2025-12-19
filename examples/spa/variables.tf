variable "website-domain-main" {
  description = "Main website domain (e.g., app.example.com)"
  type        = string
}

variable "domains-zone-root" {
  description = "Root domain for Route53 hosted zone"
  type        = string
}

variable "support-spa" {
  description = "Enable SPA support (404 redirects to index.html)"
  type        = bool
  default     = true
}
