# cycling
A package for tracking and measuring various cycling related metrics.

### ToDo
- Trends metrics
    - tests for trends metrics and new helper functions.
    - Simple Moving Avgs for trends metrics?
    - ~~CTL/Fitness~~
    - ~~ATL/Fatigue~~
    - ~~TSB/Form(freshness)~~
- Ride logic
    - Ride analyzer
    - Ride suggester
        - Consider user provided available hours/days of week.
        - Target improvement areas (based on performance adaptations logic).
        - Weigh zones based on performance adaptations.
        - Use 80/20 split (or similar logic to suggest Zone1-2 vs Zone4+ time?
    - ~~Initial ride logic~~
- Performance adaptations logic
- Figure out why NormalizedPower calculates 249 when all []ints are 250
- ~~Metrics functions~~
    - ~~Write tests for generic metrics functions.~~
    - ~~Uncouple functions from structs?~~
    - ~~Potentially break out generic functions to metrics_generics?~~
    - ~~Can we actually just remove the need to hold []int at all?~~
    - ~~Update metrics test to use ride test data instead of seperate ones.~~
    - ~~Remove calculate method from power/hr metrics.~~
    - ~~Finish Unit tests.~~
- ~~Write tests for Ride methods.~~
- ~~Tests for helper functions.~~
- ~~Zones logic~~
    - ~~Tests for HearRate zones methods~~
    - ~~Finish Zones tests.~~
    - ~~PowerZone match based on provided power~~
    - ~~Zone calculator based on % of FTP~~
    - ~~Zone standard templates (which provides %'s)~~
        - ~~Include FTP/FTHR max %s~~
        - ~~Separate file with default zone data~~.
        - ~~Method on Zones to return []float64 list of max pwr/hr~~
        - ~~Zone required recovery?~~
    - ~~Zone duration calculations?~~
    - ~~Sort method for zones slice (sort smallest zone to largest)~~
    - ~~Tests for zone sort functions~~
    - ~~Duplicate logic to HearRate zones~~
    - ~~Determine if preceived effort zones are needed.~~

