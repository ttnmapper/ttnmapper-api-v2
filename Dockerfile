# Setup environemtn by getting required packages
FROM python:3.6.8-alpine3.8
# Dump log messages immediately. Hopefully quick enough before something crashes.
ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install -r requirements.txt 

COPY . .

CMD [ "gunicorn", "-b", "0.0.0.0:8000", "ttnm_backend.wsgi"]