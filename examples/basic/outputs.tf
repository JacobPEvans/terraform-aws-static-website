output "website_cdn_root_id" {
  description = "CloudFront distribution ID for the main website"
  value       = module.static_website.website_cdn_root_id
}

output "website_root_s3_bucket" {
  description = "Name of the S3 bucket hosting the website content"
  value       = module.static_website.website_root_s3_bucket
}

output "website_logs_s3_bucket" {
  description = "Name of the S3 bucket storing access logs"
  value       = module.static_website.website_logs_s3_bucket
}

output "website_redirect_s3_bucket" {
  description = "Name of the S3 bucket handling redirects"
  value       = module.static_website.website_redirect_s3_bucket
}
