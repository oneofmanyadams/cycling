# cycling
A package for tracking and measuring various cycling related metrics.

### ToDo
- Zones logic
    - ~~PowerZone match based on provided power~~
    - ~~Zone calculator based on % of FTP~~
    - Zone duration calculations?
    - Zone standard templates (which provides %'s)
    - Duplicate logic to HearRate zones
    - Determine if prreceived effort zones are needed.
- Session analyzer
- Session suggester
- Trends metrics
- Metrics functions
    - ~~Finish Unit tests.~~
    - Remove calculate method from power/hr metrics.
    - Can we actually just remove the need to hold []int at all?
    - Potentially break out generic functions to metrics_generics?
    - Uncouple functions from structs?
    - Figure out why NormalizedPower calculates 249 when all []ints are 250
- ~~Ride logic~~
- ~~Write tests for Ride methods.~~
