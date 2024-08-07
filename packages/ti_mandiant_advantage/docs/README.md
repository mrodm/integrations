# Mandiant Advantage

## Overview

The [Mandiant Advantage](https://www.mandiant.com/advantage) integration allows users to retrieve IOCs (Indicators of Compromise) from the Threat Intelligence Advantage Module. 

These indicators can be used for correlation in Elastic Security to help discover potential threats. Mandiant Threat Intelligence gives security practitioners unparalleled visibility and expertise into threats that matter to their business right now.

Our threat intelligence is compiled by over 500 threat intelligence analysts across 30 countries, researching actors via undercover adversarial pursuits, incident forensics, malicious infrastructure reconstructions and actor identification processes that comprise the deep knowledge embedded in the Mandiant Intel Grid.

## Data streams

The Mandiant Advantage integration collects one type of data stream: `threat_intelligence`

### **Threat Intelligence**

IOCs are retrieved via the Mandiant Threat Intelligence API.


## Compatibility

- This integration has been tested against the Threat Intelligence API v4.


## Requirements

You need Elasticsearch for storing and searching your data and Kibana for visualizing and managing it.
You can use our hosted Elasticsearch Service on Elastic Cloud, which is recommended, or self-manage the Elastic Stack on your own hardware.

## Setup

For step-by-step instructions on how to set up an integration, see the
[Getting started](https://www.elastic.co/guide/en/welcome-to-elastic/current/getting-started-observability.html) guide.

For instructions on how to get Threat Intelligence API v4 credentials, see the [Mandiant Documentation Portal.](https://docs.mandiant.com/home/mati-threat-intelligence-api-v4#tag/Getting-Started)

### Filtering IOCs

The integration allows you to filter the amount of IOCs that are ingested, by using the following configuration parameters:

* **Initial interval**
  * The time in the past to start the collection of Indicator data from, based on an indicators last_update date. 
  * Supported units for this parameter are h/m/s. The default value is 720h (i.e 30 days)
  * You may reduce this interval if you do not want as much historical data to be ingested when the integration first runs.
* **Minimum IC-Score**
  * Indicators that have an IC-Score greater than or equal to the given value will be collected. 
  * Indicators with any IC-Score will be collected if a value is set to 0.
  * You might set this to a different value such as 80, to ensure that only high confidence indicators are ingested.  

## Logs reference

### Threat Intelligence

Retrieves IOCs using the Mandiant Threat Intelligence API over time.

An example event for `threat_intelligence` looks as following:

```json
{
    "@timestamp": "2023-04-25T09:36:05.822Z",
    "agent": {
        "ephemeral_id": "3cf850f4-d7a9-4302-9745-cb0d0b408c1e",
        "id": "8299ae35-ee0e-4107-9acb-1b6acfdda1fb",
        "name": "docker-fleet-agent",
        "type": "filebeat",
        "version": "8.13.0"
    },
    "data_stream": {
        "dataset": "ti_mandiant_advantage.threat_intelligence",
        "namespace": "99619",
        "type": "logs"
    },
    "ecs": {
        "version": "8.11.0"
    },
    "elastic_agent": {
        "id": "8299ae35-ee0e-4107-9acb-1b6acfdda1fb",
        "snapshot": false,
        "version": "8.13.0"
    },
    "event": {
        "agent_id_status": "verified",
        "category": [
            "threat"
        ],
        "created": "2024-08-02T05:42:35.442Z",
        "dataset": "ti_mandiant_advantage.threat_intelligence",
        "ingested": "2024-08-02T05:42:45Z",
        "kind": "enrichment",
        "module": "ti_mandiant_advantage_threat_intelligence",
        "original": "{\"first_seen\":\"2022-09-06T00:46:38.000Z\",\"id\":\"fqdn--33bf4df5-3564-51e3-84f1-ca9d5bc2329e\",\"is_publishable\":true,\"last_seen\":\"2023-03-23T21:42:34.000Z\",\"last_updated\":\"2023-04-25T09:36:05.822Z\",\"misp\":{\"akamai\":false,\"alexa\":false,\"alexa_1M\":true,\"amazon-aws\":false,\"apple\":false,\"automated-malware-analysis\":false,\"bank-website\":false,\"captive-portals\":false,\"cisco_1M\":true,\"cisco_top1000\":false,\"cisco_top10k\":false,\"cisco_top20k\":false,\"cisco_top5k\":false,\"cloudflare\":false,\"common-contact-emails\":false,\"common-ioc-false-positive\":false,\"covid\":false,\"covid-19-cyber-threat-coalition-whitelist\":false,\"covid-19-krassi-whitelist\":false,\"crl-hostname\":false,\"crl-ip\":false,\"dax30\":false,\"disposable-email\":false,\"dynamic-dns\":false,\"eicar.com\":false,\"empty-hashes\":false,\"fastly\":false,\"google\":false,\"google-chrome-crux-1million\":true,\"google-gcp\":false,\"google-gmail-sending-ips\":false,\"googlebot\":false,\"ipv6-linklocal\":false,\"majestic_million\":true,\"majestic_million_1M\":true,\"microsoft\":false,\"microsoft-attack-simulator\":false,\"microsoft-azure\":false,\"microsoft-azure-appid\":false,\"microsoft-azure-china\":false,\"microsoft-azure-germany\":false,\"microsoft-azure-us-gov\":false,\"microsoft-office365\":false,\"microsoft-office365-cn\":false,\"microsoft-office365-ip\":false,\"microsoft-win10-connection-endpoints\":false,\"moz-top500\":false,\"mozilla-CA\":false,\"mozilla-IntermediateCA\":false,\"multicast\":false,\"nioc-filehash\":false,\"ovh-cluster\":false,\"parking-domain\":false,\"parking-domain-ns\":false,\"phone_numbers\":false,\"public-dns-hostname\":false,\"public-dns-v4\":false,\"public-dns-v6\":false,\"public-ipfs-gateways\":false,\"rfc1918\":false,\"rfc3849\":false,\"rfc5735\":false,\"rfc6598\":false,\"rfc6761\":false,\"second-level-tlds\":true,\"security-provider-blogpost\":false,\"sinkholes\":false,\"smtp-receiving-ips\":false,\"smtp-sending-ips\":false,\"stackpath\":false,\"tenable-cloud-ipv4\":false,\"tenable-cloud-ipv6\":false,\"ti-falsepositives\":false,\"tlds\":true,\"tranco\":true,\"tranco10k\":true,\"university_domains\":false,\"url-shortener\":false,\"vpn-ipv4\":false,\"vpn-ipv6\":false,\"whats-my-ip\":false,\"wikimedia\":false},\"mscore\":27,\"sources\":[{\"category\":[\"test\"],\"first_seen\":\"2022-09-06T00:46:38.722+0000\",\"last_seen\":\"2023-03-23T21:42:34.707+0000\",\"osint\":true,\"source_name\":\"dtm.blackbeard\"},{\"category\":[],\"first_seen\":\"2022-11-29T16:24:52.984+0000\",\"last_seen\":\"2022-11-29T16:24:52.984+0000\",\"osint\":true,\"source_name\":\"dtm.vanellope\"}],\"type\":\"fqdn\",\"value\":\"ru.wikibooks.org\"}",
        "risk_score": 27,
        "type": [
            "indicator"
        ]
    },
    "input": {
        "type": "httpjson"
    },
    "mandiant": {
        "threat_intelligence": {
            "ioc": {
                "categories": [
                    "test"
                ],
                "first_seen": "2022-09-06T00:46:38.000Z",
                "id": "fqdn--33bf4df5-3564-51e3-84f1-ca9d5bc2329e",
                "last_seen": "2023-03-23T21:42:34.000Z",
                "last_update_date": "2023-04-25T09:36:05.822Z",
                "mscore": 27,
                "sources": [
                    {
                        "category": [
                            "test"
                        ],
                        "first_seen": "2022-09-06T00:46:38.722+0000",
                        "last_seen": "2023-03-23T21:42:34.707+0000",
                        "osint": true,
                        "source_name": "dtm.blackbeard"
                    },
                    {
                        "first_seen": "2022-11-29T16:24:52.984+0000",
                        "last_seen": "2022-11-29T16:24:52.984+0000",
                        "osint": true,
                        "source_name": "dtm.vanellope"
                    }
                ],
                "type": "fqdn",
                "value": "ru.wikibooks.org"
            }
        }
    },
    "tags": [
        "preserve_original_event",
        "forwarded",
        "mandiant-threat-intelligence-indicator"
    ],
    "threat": {
        "feed": {
            "name": "Mandiant Threat Intelligence"
        },
        "indicator": {
            "confidence": "Low",
            "first_seen": "2022-09-06T00:46:38.000Z",
            "last_seen": "2023-03-23T21:42:34.000Z",
            "marking": {
                "tlp": "GREEN",
                "tlp_version": "2.0"
            },
            "modified_at": "2023-04-25T09:36:05.822Z",
            "provider": [
                "dtm.blackbeard",
                "dtm.vanellope"
            ],
            "type": "domain-name",
            "url": {
                "domain": "ru.wikibooks.org"
            }
        }
    }
}
```

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Event timestamp. | date |
| cloud.account.id | The cloud account or organization id used to identify different entities in a multi-tenant environment. Examples: AWS account id, Google Cloud ORG Id, or other unique identifier. | keyword |
| cloud.image.id | Image ID for the cloud instance. | keyword |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| event.dataset | Event dataset | constant_keyword |
| host.os.build | OS build information. | keyword |
| host.os.codename | OS codename, if any. | keyword |
| input.type | Input type | keyword |
| log.offset | Log offset | long |
| mandiant.threat_intelligence.ioc.associated_hashes | List of associated hashes and their types. | object |
| mandiant.threat_intelligence.ioc.attributed_associations | List of attributed associations that this indicator has to other Malware families or Actors. | object |
| mandiant.threat_intelligence.ioc.categories | Categories associated with this indicator. | keyword |
| mandiant.threat_intelligence.ioc.first_seen | IOC first seen date. | date |
| mandiant.threat_intelligence.ioc.id | IOC internal ID. | keyword |
| mandiant.threat_intelligence.ioc.is_exclusive | Whether the indicator is exclusive to Mandiant or not. | boolean |
| mandiant.threat_intelligence.ioc.last_seen | IOC last seen date. | date |
| mandiant.threat_intelligence.ioc.last_update_date | IOC last update date. | date |
| mandiant.threat_intelligence.ioc.mscore | M-Score (IC-Score) between 0 - 100. | integer |
| mandiant.threat_intelligence.ioc.sources | List of the indicator sources. | object |
| mandiant.threat_intelligence.ioc.type | IOC type. | keyword |
| mandiant.threat_intelligence.ioc.value | IOC value. | keyword |
| threat.indicator.first_seen | The date and time when intelligence source first reported sighting this indicator. | date |
| threat.indicator.last_seen | The date and time when intelligence source last reported sighting this indicator. | date |
| threat.indicator.modified_at | The date and time when intelligence source last modified information for this indicator. | date |

