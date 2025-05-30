# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .card import *
from .coupon import *
from .customer import *
from .file import *
from .portal_configuration import *
from .price import *
from .product import *
from .promotion_code import *
from .provider import *
from .shipping_rate import *
from .tax_rate import *
from .webhook_endpoint import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_stripe.config as __config
    config = __config
    import pulumi_stripe.region as __region
    region = __region
else:
    config = _utilities.lazy_import('pulumi_stripe.config')
    region = _utilities.lazy_import('pulumi_stripe.region')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "stripe",
  "mod": "index/card",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/card:Card": "Card"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/coupon",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/coupon:Coupon": "Coupon"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/customer",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/customer:Customer": "Customer"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/file",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/file:File": "File"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/portalConfiguration",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/portalConfiguration:PortalConfiguration": "PortalConfiguration"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/price",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/price:Price": "Price"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/product",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/product:Product": "Product"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/promotionCode",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/promotionCode:PromotionCode": "PromotionCode"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/shippingRate",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/shippingRate:ShippingRate": "ShippingRate"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/taxRate",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/taxRate:TaxRate": "TaxRate"
  }
 },
 {
  "pkg": "stripe",
  "mod": "index/webhookEndpoint",
  "fqn": "pulumi_stripe",
  "classes": {
   "stripe:index/webhookEndpoint:WebhookEndpoint": "WebhookEndpoint"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "stripe",
  "token": "pulumi:providers:stripe",
  "fqn": "pulumi_stripe",
  "class": "Provider"
 }
]
"""
)
