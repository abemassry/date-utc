# date-utc
Convert times to UTC quickly on the command line

## Usage

If your system time is set to EDT
```
$> date-utc 1:15pm
Mon 16 Mar 2026 17:15:00 UTC
Mon 16 Mar 2026 06:15:00 PM CET
Mon 16 Mar 2026 01:15:00 PM EDT
Mon 16 Mar 2026 10:15:00 AM PDT
Mon 16 Mar 2026 10:45:00 PM IST

$> date-utc
Thu 03 Aug 2023 16:06:07 UTC

$> date-utc 10:30am
Mon 16 Mar 2026 14:30:00 UTC
Mon 16 Mar 2026 03:30:00 PM CET
Mon 16 Mar 2026 10:30:00 AM EDT
Mon 16 Mar 2026 07:30:00 AM PDT
Mon 16 Mar 2026 08:00:00 PM IST

$> date-utc 8am pdt
Mon 16 Mar 2026 15:00:00 UTC
Mon 16 Mar 2026 04:00:00 PM CET
Mon 16 Mar 2026 11:00:00 AM EDT
Mon 16 Mar 2026 08:00:00 AM PDT
Mon 16 Mar 2026 08:30:00 PM IST
```

## Install
1. clone this repo
2. `ln -s <location-cloned>/date-utc ~/bin/date-utc`
