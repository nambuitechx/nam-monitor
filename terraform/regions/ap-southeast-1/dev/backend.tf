terraform {
    backend "s3" {
        bucket = "nam-monitor-tfstate-490004621103"
        key    = "dev/terraform.tfstate"
        region = "ap-southeast-1"
    }
}