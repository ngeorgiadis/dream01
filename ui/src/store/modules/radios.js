/* eslint-disable no-console */
import axios from "axios";

const state = {
	stations: [],
	favourites: [],
	tuneBaseURL: "",
};

const getters = {
	getStations: state => state.stations,
};

const mutations = {
	setTop500: (state, payload) => {
		state.stations = payload;
	},
};

const actions = {
	getTop500: ({ commit }) => {
		axios.get("/top").then(res => {
			//this.tuneBaseURL = res.data.tune_in.base;

			let stations = res.data.stations.sort((a, b) => {
				return a.lc > b.lc;
			});

			commit("setTop500", stations);
		});
	},
};

export default {
	state,
	getters,
	mutations,
	actions,
};
