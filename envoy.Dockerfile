FROM envoyproxy/envoy-dev:11627cb0bd1d77b18382205682a78d245c132b55
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml