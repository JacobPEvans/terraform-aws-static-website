output "logs_bucket_id" {
  description = "ID of the logs S3 bucket"
  value       = aws_s3_bucket.website_logs.id
}

output "logs_bucket_name" {
  description = "Name of the logs S3 bucket"
  value       = aws_s3_bucket.website_logs.bucket
}

output "logs_bucket_domain_name" {
  description = "Domain name of the logs S3 bucket"
  value       = aws_s3_bucket.website_logs.bucket_domain_name
}

output "root_bucket_id" {
  description = "ID of the root S3 bucket"
  value       = aws_s3_bucket.website_root.id
}

output "root_bucket_name" {
  description = "Name of the root S3 bucket"
  value       = aws_s3_bucket.website_root.bucket
}

output "root_bucket_arn" {
  description = "ARN of the root S3 bucket"
  value       = aws_s3_bucket.website_root.arn
}

output "root_bucket_website_endpoint" {
  description = "Website endpoint of the root S3 bucket"
  value       = aws_s3_bucket_website_configuration.website_root.website_endpoint
}

output "redirect_bucket_id" {
  description = "ID of the redirect S3 bucket"
  value       = aws_s3_bucket.website_redirect.id
}

output "redirect_bucket_name" {
  description = "Name of the redirect S3 bucket"
  value       = aws_s3_bucket.website_redirect.bucket
}

output "redirect_bucket_website_endpoint" {
  description = "Website endpoint of the redirect S3 bucket"
  value       = aws_s3_bucket_website_configuration.website_redirect.website_endpoint
}
