variable "ibmcloud_api_key" {
}

variable "cos_service_name" {
  default = "yus-ocp"
}

variable "cos_service_plan" {
  default = "standard"
}

variable "cluster_node_flavor" {
  default = "bx2.16x64"
}

variable "cluster_kube_version" {
  default = "4.11_openshift"
}

variable "deafult_worker_pool_count" {
  default = "1"
}

variable "region" {
  default = "us-south"
}

variable "resource_group" {
  default = "ibm-satellite-dev"
}

variable "cluster_name" {
  default = "yus-roks-on-vpc"
}

variable "worker_pool_name" {
  default = "workerpool"
}

variable "entitlement" {
  default = ""
}
