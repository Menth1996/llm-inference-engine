
# LLM Inference Engine

This project provides a high-performance inference engine specifically designed for serving quantized Large Language Models (LLMs). It focuses on low-latency and high-throughput inference.

## Features
- Optimized for quantized models
- Low-latency inference
- High-throughput serving
- Scalable architecture

## Getting Started

```bash
# Clone the repository
git clone https://github.com/Menth1996/llm-inference-engine.git
cd llm-inference-engine

# Build the engine
go build -o llm-inference-engine ./cmd/server

# Run the server
./llm-inference-engine
```

## Technologies Used
- Go
- ONNX Runtime
- Quantization techniques

## License
This project is licensed under the MIT License.
