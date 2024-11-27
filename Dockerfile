FROM python:3.12.7-alpine3.20

WORKDIR /app

# Copy all the files from the src directory to the container
COPY ./app/ ./
COPY requirements.txt ./

# Install dependencies
RUN  pip install --no-cache-dir  --requirement requirements.txt

CMD ["fastapi", "run", "--port", "3000", "main.py"] 

