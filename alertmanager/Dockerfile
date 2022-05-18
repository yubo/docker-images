FROM quay.io/prometheus/alertmanager:v0.24.0 as src

FROM ybbbbasdf/debuger:latest
COPY --from=src /bin/alertmanager /bin/alertmanager
COPY --from=src /etc/alertmanager/alertmanager.yml /etc/alertmanager/alertmanager.yml
WORKDIR /alertmanager
ENTRYPOINT ["/bin/alertmanager"]
CMD ["--config.file=/etc/alertmanager/alertmanager.yml", "--storage.path=/alertmanager"]
