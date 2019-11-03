<template>
	<div class="home">
		<md-card
			md-with-hover
			v-for="r in radios"
			:key="r.id"
			:class="{ active: r.id == selectedStation }"
			:title="r.name"
		>
			<md-card-header>
				<md-card-header-text>
					<div class="md-title">{{ r.name | normalize }}</div>
					<div class="md-subhead">{{ r.genre }} ({{ r.lc }})</div>
				</md-card-header-text>
				<md-icon class="md-accent md-size-2x">{{
					r.id == selectedStation ? "speaker" : ""
				}}</md-icon>
			</md-card-header>
			<md-card-content>
				<md-card-media class="thumb">
					<img
						v-show="r.logo.length > 0"
						:src="getLogo(r)"
						alt="logo"
					/>

					<md-icon v-if="getLogo(r).length == 0" class="md-size-3x"
						>radio</md-icon
					>
				</md-card-media>
			</md-card-content>
			<md-card-actions>
				<md-button @click="tune(r.id)">Tune</md-button>
			</md-card-actions>
		</md-card>
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
			radios: []
		};
	},

	created() {
		axios.get("/top").then(res => {
			this.tuneBaseURL = res.data.tune_in.base;

			this.radios = res.data.stations.sort((a, b) => {
				return a.lc > b.lc;
			});
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
				id: id
			});
		}
	},

	filters: {
		normalize(s) {
			const regex = /(\w|\.|,){18,}/gm;
			var res = s.replace(regex, "_");

			if (res.length >= 30) {
				return res.substring(0, 29) + "_";
			}
			return res;
		}
	}
};
</script>

<style lang="scss" scoped>
.home {
	margin: 10px;

	display: grid;
	grid-column-gap: 10px;
	grid-row-gap: 10px;
}

.thumb img {
	max-width: 150px;
	max-height: 150px;
	width: unset !important;
}

.md-card-content {
	height: 150px;
}

.md-title {
	overflow: hidden;
	text-overflow: ellipsis;
}

@media only screen and (max-width: 600px) {
	.home {
		display: grid;
		grid-template-columns: 1fr;
	}
}
@media only screen and (min-width: 600px) {
	.home {
		display: grid;
		grid-template-columns: 1fr 1fr;
	}
}
@media only screen and (min-width: 901px) {
	.home {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
	}
}
</style>

