<template>
	<div class="home">
		<div class="stations">
			<v-card v-for="r in getFavourites" :key="r.id" @click="tune(r)">
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
import { mapActions, mapGetters } from "vuex";
import axios from "axios";
export default {
	data() {
		return {
			favourites: [],
		};
	},

	created() {
		axios.get("/fav").then(res => {
			this.favourites = res.data;
		});
	},

	computed: {
		...mapGetters(["getStations", "getCurrentStation"]),

		getFavourites() {
			return this.favourites;
		},

		selectedStation() {
			return this.getCurrentStation.id;
		},
	},

	methods: {
		...mapActions(["getTop500", "tuneFavourite"]),

		getLogo(r) {
			return r.logo.trim();
		},

		tune(r) {
			this.tuneFavourite(r);
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

<style lang="scss">
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
