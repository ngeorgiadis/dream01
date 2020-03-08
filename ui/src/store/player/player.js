/* eslint-disable no-console */
import axios from "axios";

const state = {
	playerStatus: "paused",
	source: "http://dream01.gr:8000/stream",
};

const getters = {
	getPlayerStatus: state => {
		return state.playerStatus;
	},

	getPlayerSource: state => {
		return state.source;
	},
};

const mutations = {
	setPlayerStatus: (state, payload) => {
		state.playerStatus = payload;
	},

	setPlayerSource: (state, payload) => {
		state.source = payload;
	},
};

const actions = {
	setPlayerStatus: ({ commit }, payload) => {
		commit("setPlayerStatus", payload);
	},

	setPlayerSource: ({ commit }, payload) => {
		commit("setPlayerSource", payload);
	},

	tuneStationByID: ({ dispatch }, payload) => {
		//http://yp.shoutcast.com

		var url = `/source/${payload.id}`;
		axios.get(url, { crossdomain: true }).then(res => {
			dispatch("setPlayerSource", res.data);
		});
	},
};

export default {
	state,
	getters,
	mutations,
	actions,
};
