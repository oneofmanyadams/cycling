# cycling
A package for tracking and measuring various cycling related metrics.

### ToDo
- Zones logic
    - ~~PowerZone match based on provided power~~
    - ~~Zone calculator based on % of FTP~~
    - ~~Zone standard templates (which provides %'s)~~
        - ~~Include FTP/FTHR max %s~~
        - ~~Separate file with default zone data~~.
        - ~~Method on Zones to return []float64 list of max pwr/hr~~
        - ~~Zone required recovery?~~
    - ~~Zone duration calculations?~~
    - ~~Sort method for zones slice (sort smallest zone to largest)~~
    - Tests for zone sort functions
    - Duplicate logic to HearRate zones
    - ~~Determine if preceived effort zones are needed.~~
- Ride logic
    - ~~Initial ride logic~~
    - Ride analyzer
    - Ride suggester
        - Consider user provided available hours/days of week.
        - Target improvement areas (based on performance adaptations logic).
        - Weigh zones based on performance adaptations.
        - Use 80/20 split (or similar logic0 to suggest Zone1-2 vs Zone4+ time?
- Trends metrics
    - CTL/Fitness
    - ATL/Fatigue
    - TSB/Form(freshness)
- Performance adaptations logic
- Metrics functions
    - ~~Finish Unit tests.~~
    - Add additional metrics
    - Remove calculate method from power/hr metrics.
    - Can we actually just remove the need to hold []int at all?
    - Potentially break out generic functions to metrics_generics?
    - Uncouple functions from structs?
    - Figure out why NormalizedPower calculates 249 when all []ints are 250
- ~~Write tests for Ride methods.~~
- ~~Tests for helper functions.~~
