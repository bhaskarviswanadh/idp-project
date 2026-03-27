# Use the official Python image
FROM python:3.9-slim

# Set the working directory
WORKDIR /app

# Install flask
RUN pip install flask

# Copy the application code
COPY app.py .

# Expose the port
EXPOSE 5000

# Run the application
CMD ["python", "app.py"]
