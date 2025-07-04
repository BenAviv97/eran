<!-- /home/${USER}/livestreampro/docs/sre/runbooks/stream-outage.md -->
# Stream Outage

Use this runbook when live streams are unavailable or failing to load for viewers.

## Detection
- Monitor alerts from the streaming pipeline (e.g., CDN 5xx errors).
- Reports from on-call engineers or support.

## Immediate Actions
1. Confirm the scope of outage and affected regions.
2. Restart impacted streaming services or pods if necessary.
3. Notify stakeholders in the incident channel.

## Resolution
- Investigate root cause (e.g., transcoder failure, network issues).
- Apply fixes or rollbacks as required.

## Postmortem
- Document timeline, impact, and resolution steps.
- Identify follow-up actions to prevent recurrence.
