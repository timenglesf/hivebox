"""
Module for getting the current version of the application.
"""

from os import getenv
from pydantic import BaseModel


class Version(BaseModel):
    version: str


def get_env_version():
    """
    Get the version of the application from the environment.
    """
    version = getenv("VERSION")
    if version is None:
        raise ValueError("VERSION environment variable is not set")
    return version


def get_version_json() -> Version:
    """
    Get the version of the application as a JSON object.
    """
    return Version(version=get_env_version())
