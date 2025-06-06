const env = process.env.NODE_ENV;

export const BASE_WEBAPP_URL = "http://localhost:3001";
export const BASE_URL = process.env.NEXT_PUBLIC_BASE_URL + "/api"
export const BASE_WS_URL = process.env.NEXT_PUBLIC_BASE_WS_URL + "/api"

export const localConstants = {
  ROOM_NAME_KEY: "roomName",
  USERNAME_KEY: "username",
};

export const cookieConstants = {
  USER_KEY: "user",
};

export const ROOM_NAME_INPUT_LIMIT = 30;
export const NAME_INPUT_LIMIT = 12;

export const messages = {
  ROOM_JOINED: "room-joined",
  VOTE_SUBMITTED: "vote-submitted",
  REVEAL_CARDS: "reveal-cards",
  CARDS_REVEALED: "cards-revealed",
  START_NEW_VOTING: "start-new-voting",
  NEW_VOTING_STARTED: "new-voting-started",
  IS_AFK: "is-afk",
  DISCONNECTED: "disconnected",
  KICK_USER: "kick",
  USER_KICKED: "client-kicked",
};

export const gameStates = {
  IN_PROGRESS: "IN_PROGRESS",
  CARDS_REVEALED: "CARDS_REVEALED",
};

export const voteCardValues = {
  CONFUSED: -1, // the '?' card
  NOT_SELECTED: -2,
  PRIVATE: -3, // when card is selected but not revealed, the value is private
  EMPTY: -4, // used on the client side only, when restarting the voting session
};
