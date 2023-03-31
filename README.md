# Get YouTube Transcript

Get YouTube Transcript is an open-source project that provides a simple Golang server for extracting transcripts from YouTube videos. The server can be deployed on fly.io and as a Docker container. It exposes a single endpoint that accepts a YouTube video ID and returns the video transcript in response.

## Features

- Extracts transcripts from YouTube videos
- Supports multiple languages

## Getting Started

Clone the repository:

```bash
git clone https://github.com/aiocean/get-youtube-transcript.git
cd get-youtube-transcript
```

Build the server:

```bash
go build -o get-youtube-transcript
```

Run the server:

```bash
./get-youtube-transcript --port 8080
```

The server will start on port 8080 by default. You can now send requests to the `/transcripts` endpoint:


## Usage

The server exposes a single endpoint that accepts a YouTube video ID and returns the video transcript in response.

```bash
curl -X GET "http://localhost:8080/transcripts/YOUR_VIDEO_ID"
```

Replace YOUR_VIDEO_ID with the ID of the YouTube video you want to extract the transcript from.


## Deploying to Fly.io

1. Install the Fly.io CLI and log in:

```bash
curl -L https://fly.io/install.sh | sh
fly login
```

2. Create a new Fly.io app:

```bash
flyctl init
```

3. Deploy the app:

```bash
flyctl deploy
```

## Deploying as a Docker container

1. Build the Docker image:

```bash
docker build -t get-youtube-transcript .
```

2. Run the container:

```bash
docker run -p 8080:8080 get-youtube-transcript
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

We welcome contributions! If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch with a descriptive name.
3. Implement your changes or additions.
4. Test your