# LLM Inference Engine

## A High-Performance Go-based Inference Engine for Serving Quantized LLMs

![Go LLM Inference](https://miro.medium.com/v2/resize:fit:1400/1*b13f_e_7i5f5s7_j5l_s_g.png)

This repository presents a high-performance, low-latency inference engine built in Go, specifically designed for serving quantized Large Language Models (LLMs). Leveraging Go's concurrency primitives and efficient memory management, this engine aims to provide a robust and scalable solution for deploying LLMs in production environments, with a focus on edge devices and resource-constrained settings.

## Features

- **Quantized Model Support**: Optimized for running INT8, FP16, and other quantized LLMs for reduced memory footprint and faster inference.
- **Concurrent Request Handling**: Utilizes Go routines and channels for efficient parallel processing of multiple inference requests.
- **Low Latency**: Designed for minimal inference latency, crucial for real-time AI applications.
- **Memory Efficiency**: Go's garbage collection and efficient data structures ensure optimal memory usage.
- **API Server**: Built-in HTTP/gRPC server for easy integration with client applications.
- **Cross-Platform**: Compiles to a single binary, making deployment straightforward across various operating systems.

## Installation

To build and run the LLM Inference Engine, you need Go (version 1.18 or higher) installed. Clone the repository and build the executable:

```bash
git clone https://github.com/Menth1996/llm-inference-engine.git
cd llm-inference-engine
go mod tidy
go build -o bin/llm-inference-engine cmd/server/main.go
```

## Usage

### Running the Inference Server

```bash
./bin/llm-inference-engine --config configs/config.yaml
```

### Example: Making an Inference Request (using curl)

```bash
curl -X POST http://localhost:8080/predict \
     -H "Content-Type: application/json" \
     -d '{"prompt": "What is the capital of France?"}'
```

## Project Structure

```
llm-inference-engine/
├── cmd/
│   └── server/
│       └── main.go             # Main entry point for the inference server
├── pkg/
│   ├── inference/
│   │   └── inference.go        # Core inference logic and model loading
│   └── model/
│       └── model.go            # LLM model definition and quantization handling
├── configs/
│   └── config.yaml             # Configuration file for model paths, ports, etc.
├── tests/
│   └── inference_test.go       # Unit tests for inference logic
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── LICENSE                     # MIT License
└── README.md                   # Project overview and documentation
```

## Contributing

We welcome contributions! Please see `CONTRIBUTING.md` for guidelines on how to submit pull requests, report bugs, and suggest new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
