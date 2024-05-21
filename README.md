# cycling
A package for tracking and measuring various cycling related metrics.

### ToDo
- Zones logic
    - ~~PowerZone match based on provided power~~
    - ~~Zone calculator based on % of FTP~~
    - Zone standard templates (which provides %'s)
        - Populate zone template via json with actual data in configs dir.
    - Zone duration calculations?
    - Sort method for zones slice (sort smallest zone to largest)
    - Duplicate logic to HearRate zones
    - Determine if preceived effort zones are needed.
- Ride logic
    - ~~Initial ride logic~~
    - Session analyzer
    - Session suggester
        - Consider user provided available hours/days of week.
        - Target improvement areas (based on performance adaptations logic).
        - Weigh zones based on performance adaptations.
- Trends metrics
- Performance adaptations logic
- Metrics functions
    - ~~Finish Unit tests.~~
    - Remove calculate method from power/hr metrics.
    - Can we actually just remove the need to hold []int at all?
    - Potentially break out generic functions to metrics_generics?
    - Uncouple functions from structs?
    - Figure out why NormalizedPower calculates 249 when all []ints are 250
- ~~Write tests for Ride methods.~~
- ~~Tests for helper functions.~~
