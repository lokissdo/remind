# Ref to https://gitlab.com/nvidia/container-images/cuda/blob/master/dist/12.5.1/ubuntu2404/runtime/Dockerfile to get the latest base image

FROM nvidia/cuda:12.5.1-base-ubuntu24.04

# Package version control
ARG PYTHON_VERSION=3.12.3
ARG CUDA_VERSION=11.7
ARG PYTORCH_VERSION=2.3.1
ARG CUDA_CHANNEL=nvidia
ARG INSTALL_CHANNEL=pytorch
ARG CLIP_VERSION=ViT-L/14@336px
ARG BERT_VERSION=all-mpnet-base-v2

# Setup workdir and non-root user
ARG USERNAME=thanh.nguyen
WORKDIR /home/$USERNAME/workspace/

ENV TZ=Asia/Ho_Chi_Minh \
    DEBIAN_FRONTEND=noninteractive \
    PATH=/opt/conda/bin:$PATH \
    NVIDIA_VISIBLE_DEVICES=all \
    NVIDIA_DRIVER_CAPABILITIES=compute,utility \
    LD_LIBRARY_PATH=/usr/local/nvidia/lib:/usr/local/nvidia/lib64 \
    PYTORCH_VERSION=${PYTORCH_VERSION}

# Install basic packages and create user
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl git sudo && \
    useradd --create-home --shell /bin/bash $USERNAME && \
    echo "$USERNAME ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME && \
    rm -rf /var/lib/apt/lists/*

# Install development packages
RUN --mount=type=cache,id=apt-dev,target=/var/cache/apt \
    apt-get -qq update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    ccache \
    cmake \
    gcc \
    tmux \
    libjpeg-dev \
    unzip bzip2 ffmpeg libsm6 libxext6 \
    libpng-dev && \
    rm -rf /var/lib/apt/lists/*

# Install mamba
RUN curl -fsSL -o ~/mambaforge.sh https://github.com/conda-forge/miniforge/releases/download/24.3.0-0/Mambaforge-24.3.0-0-Linux-aarch64.sh && \
    chmod +x ~/mambaforge.sh && \
    ~/mambaforge.sh -b -p /opt/conda && \
    chown -R $USERNAME:$USERNAME /opt/conda/ && \
    rm ~/mambaforge.sh

# Change ownership of home and conda directories
RUN chown -R $USERNAME:$USERNAME /home/$USERNAME/ && \
    chown -R $USERNAME:$USERNAME /opt/conda/

# Install Python and PyTorch packages
RUN /opt/conda/bin/mamba install -c "${INSTALL_CHANNEL}" -c "${CUDA_CHANNEL}" -y \
    python=${PYTHON_VERSION} \
    pytorch=${PYTORCH_VERSION} \
    torchvision \
    cudatoolkit=${CUDA_VERSION} && \
    /opt/conda/bin/mamba clean -ya

# Switch to non-root user
USER $USERNAME

# Initialize conda for bash
RUN conda init bash

# Install CLIP and other dependencies
RUN pip install --no-cache-dir git+https://github.com/openai/CLIP.git

# Copy and install additional Python dependencies
COPY ./requirements.txt $WORKDIR
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code
COPY . $WORKDIR

# Load models at runtime instead of during the build phase
CMD ["bash", "-c", "python -c \"import clip; clip.load('$CLIP_VERSION', device='cpu')\" && python -c \"from sentence_transformers import SentenceTransformer; model = SentenceTransformer('$BERT_VERSION')\" && exec \"$@\"", "--"]