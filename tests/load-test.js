import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '2m', target: 1000}, // ramp up to 10000 users over 2 minutes
    { duration: '1m', target: 1500}, // stay at 10000 users for 5 minutes
    { duration: '2m', target: 0 }, // ramp down to 0 users
  ],
};

export default function () {
  const BASE_URL = 'http://localhost:8080';

  // 1. Create a new paste
  const createPayload = JSON.stringify({
    content: `This is a test paste from k6. number${Math.random()}`,
  });
  const createParams = {
    headers: {
      'Content-Type': 'application/json',
    },
  };
  const createRes = http.post(`${BASE_URL}/pastes`, createPayload, createParams);
  check(createRes, {
    'create paste status is 201': (r) => r.status === 201,
    'create paste response has id': (r) => r.json().hasOwnProperty('ID'),
  });

  const pasteId = createRes.json().ID;

  sleep(1);

  // 2. Get all pastes
  const getPastesRes = http.get(`${BASE_URL}/pastes`);
  check(getPastesRes, {
    'get pastes status is 200': (r) => r.status === 200,
  });

  sleep(1);

  // 3. Get a specific paste
  if (pasteId) {
    const getPasteRes = http.get(`${BASE_URL}/pastes/${pasteId}`);
    check(getPasteRes, {
      'get paste status is 200': (r) => r.status === 200,
    });
  }

  sleep(1);

  
}
