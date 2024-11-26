"""
Test module for FastAPI application.
"""

from os import getenv
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)


def test_version():
    """
    Test the /version endpoint.
    """
    response = client.get("/version")
    version = getenv("VERSION")
    assert version is not None
    assert response.status_code == 200
    assert response.json() == {"version": version}
