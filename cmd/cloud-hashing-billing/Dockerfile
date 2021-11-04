FROM jrei/systemd-centos:7

COPY CloudHashingBilling.viper.yaml /etc/CloudHashingBilling/CloudHashingBilling.viper.yaml
COPY cloud-hashing-billing /usr/local/bin/cloud-hashing-billing

CMD ["cloud-hashing-billing", "run"]
