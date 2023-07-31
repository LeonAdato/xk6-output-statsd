// Example taken from k6 website, https://k6.io/docs/using-k6/scenarios/executors/ramping-vus/.
import http from 'k6/http';

export const options = {
    discardResponseBodies: true,

    scenarios: {
        contacts: {
            executor: 'ramping-arrival-rate',

            // Start iterations per `timeUnit`
            startRate: 300,

            // Start `startRate` iterations per minute
            timeUnit: '1m',

            // Pre-allocate necessary VUs.
            preAllocatedVUs: 50,


            stages: [
                // Start 300 iterations per `timeUnit` for the first minute.
                { target: 300, duration: '1m' },

                // Linearly ramp-up to starting 600 iterations per `timeUnit` over the following two minutes.
                { target: 600, duration: '2m' },

                // Continue starting 600 iterations per `timeUnit` for the following four minutes.
                { target: 600, duration: '4m' },

                // Linearly ramp-down to starting 60 iterations per `timeUnit` over the last two minutes.
                { target: 60, duration: '2m' },
            ],
        },
    },
};

export default function () {
    http.get('https://test.k6.io/contacts.php');
}
