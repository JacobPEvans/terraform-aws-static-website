output "root_record_fqdn" {
  description = "FQDN of the root DNS record"
  value       = aws_route53_record.website_cdn_root_record.fqdn
}

output "redirect_record_fqdn" {
  description = "FQDN of the redirect DNS record"
  value       = var.redirect_enabled ? aws_route53_record.website_cdn_redirect_record[0].fqdn : ""
}
