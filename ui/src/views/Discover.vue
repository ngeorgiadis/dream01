<template>
	<div class="home">
		<div class="stations">
			<v-card v-for="r in radios" :key="r.id" @click="tune(r.id)">
				<v-card-title
					:class="
						selectedStation == r.id
							? 'red darken-1 white--text'
							: 'grey lighten-4'
					"
				>
					{{ r.name | normalize }}
				</v-card-title>
				<v-card-subtitle
					:class="
						selectedStation == r.id
							? 'red darken-1'
							: 'grey lighten-4'
					"
					>{{ r.genre }} ({{ r.lc }})</v-card-subtitle
				>
				<v-card-text class="mt-4" height="100px">
					<img
						class="thumb"
						v-show="r.logo.length > 0"
						:src="getLogo(r)"
						alt="logo"
					/>
					<v-icon v-if="getLogo(r).length == 0" x-large>radio</v-icon>
				</v-card-text>
			</v-card>
		</div>
	</div>
</template>

<script>
import axios from "axios";
import { mapActions } from "vuex";

export default {
	name: "home",
	components: {},
	data() {
		return {
			selectedStation: -1,
			tuneBaseURL: "",
			radios: [],
		};
	},

	created() {
		axios.get("/top").then(res => {
			this.tuneBaseURL = res.data.tune_in.base;

			this.radios = res.data.stations
				.sort((a, b) => {
					return a.lc > b.lc;
				})
				.slice(0, 20);
		});
	},

	methods: {
		...mapActions(["tuneStationByID"]),

		getLogo(r) {
			return r.logo.trim();
		},

		tune(id) {
			this.selectedStation = id;

			this.tuneStationByID({
				base: this.tuneBaseURL,
				id: id,
			});
		},
	},

	filters: {
		normalize(s) {
			const regex = /(\w|\.|,){18,}/gm;
			var res = s.replace(regex, "_");

			if (res.length >= 30) {
				return res.substring(0, 29) + "_";
			}
			return res;
		},
	},
};
</script>

<style lang="scss" scoped>
.stations {
	overflow: auto;

	margin: 10px;
	padding: 15px;

	display: grid;
	grid-column-gap: 10px;
	grid-row-gap: 10px;
}

.thumb,
.thumb img {
	max-width: 80px;
	max-height: 80px;
	width: unset !important;
}

.md-title {
	overflow: hidden;
	text-overflow: ellipsis;
}

@media only screen and (max-width: 600px) {
	.stations {
		display: grid;
		grid-template-columns: 1fr;
	}

	.filter {
		grid-column: 1/2;
	}
}
@media only screen and (min-width: 600px) {
	.stations {
		display: grid;
		grid-template-columns: 1fr 1fr;
	}

	.filter {
		grid-column: 1/3;
	}
}
@media only screen and (min-width: 901px) {
	.stations {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
	}

	.filter {
		grid-column: 1/4;
	}
}
</style>
