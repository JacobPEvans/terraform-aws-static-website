output "root_distribution_id" {
  description = "ID of the root CloudFront distribution"
  value       = aws_cloudfront_distribution.website_cdn_root.id
}

output "root_distribution_domain_name" {
  description = "Domain name of the root CloudFront distribution"
  value       = aws_cloudfront_distribution.website_cdn_root.domain_name
}

output "root_distribution_hosted_zone_id" {
  description = "Hosted zone ID of the root CloudFront distribution"
  value       = aws_cloudfront_distribution.website_cdn_root.hosted_zone_id
}

output "redirect_distribution_id" {
  description = "ID of the redirect CloudFront distribution"
  value       = var.redirect_enabled ? aws_cloudfront_distribution.website_cdn_redirect[0].id : ""
}

output "redirect_distribution_domain_name" {
  description = "Domain name of the redirect CloudFront distribution"
  value       = var.redirect_enabled ? aws_cloudfront_distribution.website_cdn_redirect[0].domain_name : ""
}

output "redirect_distribution_hosted_zone_id" {
  description = "Hosted zone ID of the redirect CloudFront distribution"
  value       = var.redirect_enabled ? aws_cloudfront_distribution.website_cdn_redirect[0].hosted_zone_id : ""
}
