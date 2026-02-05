output "website_cdn_root_id" {
  description = "Main CloudFront Distribution ID"
  value       = module.cloudfront.root_distribution_id
}

output "website_root_s3_bucket" {
  description = "The website root bucket where resources are uploaded"
  value       = module.s3_buckets.root_bucket_name
}

output "website_logs_s3_bucket" {
  description = "The s3 bucket of the website logs"
  value       = module.s3_buckets.logs_bucket_name
}

output "website_redirect_s3_bucket" {
  description = "The s3 bucket of the website redirect bucket"
  value       = module.s3_buckets.redirect_bucket_name
}
