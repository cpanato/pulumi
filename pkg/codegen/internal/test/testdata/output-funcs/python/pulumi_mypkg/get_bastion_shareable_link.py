# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = [
    'GetBastionShareableLinkResult',
    'AwaitableGetBastionShareableLinkResult',
    'get_bastion_shareable_link',
    'get_bastion_shareable_link_output',
]

@pulumi.output_type
class GetBastionShareableLinkResult:
    """
    Response for all the Bastion Shareable Link endpoints.
    """
    def __init__(__self__, next_link=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> Optional[str]:
        """
        The URL to get the next set of results.
        """
        return pulumi.get(self, "next_link")


class AwaitableGetBastionShareableLinkResult(GetBastionShareableLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBastionShareableLinkResult(
            next_link=self.next_link)


def get_bastion_shareable_link(bastion_host_name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               vms: Optional[Sequence[pulumi.InputType['BastionShareableLink']]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBastionShareableLinkResult:
    """
    Response for all the Bastion Shareable Link endpoints.
    API Version: 2020-11-01.


    :param str bastion_host_name: The name of the Bastion Host.
    :param str resource_group_name: The name of the resource group.
    :param Sequence[pulumi.InputType['BastionShareableLink']] vms: List of VM references.
    """
    __args__ = dict()
    __args__['bastionHostName'] = bastion_host_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['vms'] = vms
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mypkg::getBastionShareableLink', __args__, opts=opts, typ=GetBastionShareableLinkResult).value

    return AwaitableGetBastionShareableLinkResult(
        next_link=__ret__.next_link)


@_utilities.lift_output_func(get_bastion_shareable_link)
def get_bastion_shareable_link_output(bastion_host_name: Optional[pulumi.Input[str]] = None,
                                      resource_group_name: Optional[pulumi.Input[str]] = None,
                                      vms: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['BastionShareableLink']]]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBastionShareableLinkResult]:
    """
    Response for all the Bastion Shareable Link endpoints.
    API Version: 2020-11-01.


    :param str bastion_host_name: The name of the Bastion Host.
    :param str resource_group_name: The name of the resource group.
    :param Sequence[pulumi.InputType['BastionShareableLink']] vms: List of VM references.
    """
    ...
