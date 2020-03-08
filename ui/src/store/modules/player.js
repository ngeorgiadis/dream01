import axios from "axios";

const state = {
	playerStatus: "paused",
	source: "http://dream01.gr:8000/stream",
	currentStation: {},
};

const getters = {
	getPlayerStatus: state => {
		return state.playerStatus;
	},

	getPlayerSource: state => {
		return state.source;
	},

	getCurrentStation: state => state.currentStation,
};

const mutations = {
	setPlayerStatus: (state, payload) => {
		state.playerStatus = payload;
	},

	setPlayerSource: (state, payload) => {
		state.source = payload;
	},

	setCurrentStation: (state, payload) => {
		state.currentStation = payload;
	},
};

const actions = {
	setPlayerStatus: ({ commit }, payload) => {
		commit("setPlayerStatus", payload);
	},

	setPlayerSource: ({ commit }, payload) => {
		commit("setPlayerSource", payload);
	},

	setCurrentStation: ({ commit }, payload) => {
		commit("setCurrentStation", payload);
	},

	tuneStation: ({ dispatch }, station) => {
		var url = `/source/${station.id}`;
		axios.get(url).then(res => {
			dispatch("setCurrentStation", station);
			dispatch("setPlayerSource", res.data[0]);
		});
	},
};

export default {
	state,
	getters,
	mutations,
	actions,
};
