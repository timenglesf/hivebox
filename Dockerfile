FROM python:3.12.7-alpine3.20

WORKDIR /app

# Copy all the files from the src directory to the container
COPY ./src/ ./
COPY requirements.txt ./

# Install dependencies
RUN pyton3 -m pip install --upgrade pip && \ 
  pip install --no-cache-dir  -r requirements.txt

CMD ["fastapi", "run", "--port", "3000", "main.py"] 

