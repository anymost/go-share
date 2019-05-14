FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./server"]