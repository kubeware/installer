#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

KUBEWARE_CATALOG_TAG=${KUBEWARE_CATALOG_TAG:-master}

crd-importer \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_elasticsearchbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_kafkabindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_mariadbbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_memcachedbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_mongodbbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_mysqlbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_perconaxtradbbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_pgbouncerbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_postgresbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_proxysqlbindings.yaml \
    --input=https://github.com/kubeware/catalog/raw/${KUBEWARE_CATALOG_TAG}/config/crd/bases/catalog.appscode.com_redisbindings.yaml \
    --out=./charts/catalog-manager/crds

crd-importer \
    --input=https://github.com/kubeware/kubeware/raw/master/deploy/crd/kubeware.dev_apiservicebindings.yaml \
    --input=https://github.com/kubeware/kubeware/raw/master/deploy/crd/kubeware.dev_apiserviceexportrequests.yaml \
    --input=https://github.com/kubeware/kubeware/raw/master/deploy/crd/kubeware.dev_apiserviceexports.yaml \
    --input=https://github.com/kubeware/kubeware/raw/master/deploy/crd/kubeware.dev_apiservicenamespaces.yaml \
    --input=https://github.com/kubeware/kubeware/raw/master/deploy/crd/kubeware.dev_clusterbindings.yaml \
    --out=./charts/kubeware-provider/crds
