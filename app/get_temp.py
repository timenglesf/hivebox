"""
Module for getting the average temperature of all sensebox sensors.
"""


def get_sensor_data():
    """
    Get the sensor data from the sensebox API.
    """
    return None


def filter_seonsor_data(sensor_data):
    """
    Filter the sensor data to only include temperature values that are less than an hour old.
    """

    return [
        data["value"] for data in sensor_data if data["value_type"] == "temperature"
    ]


def calculate_average_temperature(sensor_data):
    """
    Calculate the average temperature of all sensebox sensors.
    """
    return sum(sensor_data) / len(sensor_data)


def get_average_tempeature():
    """
    Get the average temperature of all sensebox sensors.
    """
    return None
