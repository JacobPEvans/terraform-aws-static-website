output "website_cdn_root_id" {
  description = "CloudFront distribution ID for the main website"
  value       = module.static_website_with_lambda.website_cdn_root_id
}

output "website_root_s3_bucket" {
  description = "Name of the S3 bucket hosting the website content"
  value       = module.static_website_with_lambda.website_root_s3_bucket
}

output "website_logs_s3_bucket" {
  description = "Name of the S3 bucket storing access logs"
  value       = module.static_website_with_lambda.website_logs_s3_bucket
}

output "website_redirect_s3_bucket" {
  description = "Name of the S3 bucket handling redirects"
  value       = module.static_website_with_lambda.website_redirect_s3_bucket
}

output "lambda_edge_function_arn" {
  description = "ARN of the Lambda@Edge function"
  value       = aws_lambda_function.security_headers.qualified_arn
}
