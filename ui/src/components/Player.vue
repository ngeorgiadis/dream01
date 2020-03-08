<template>
	<div>
		<v-btn
			@click="tooglePlay"
			class="md-icon-button md-theme-footer md-primary"
		>
			<v-icon>{{ playerIcon }}</v-icon>
		</v-btn>
		<v-btn class="md-icon-button md-theme-footer md-primary">
			<v-icon>volume_down</v-icon>
		</v-btn>
		<v-btn class="md-icon-button md-theme-footer md-primary">
			<v-icon>volume_up</v-icon>
		</v-btn>
		{{ currentStationInfo }}
		<audio id="player" :src="''"></audio>
	</div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
	data() {
		return {
			player: {},
			playerStatus: "paused",
			playerIcon: "play_arrow",
		};
	},

	props: {
		source: {
			type: String,
			default: "http://dream01.gr:8000/stream",
		},
	},

	created() {},

	computed: {
		...mapGetters([
			"getPlayerStatus",
			"getPlayerSource",
			"getCurrentStation",
		]),

		currentStationInfo() {
			let c = this.getCurrentStation;
			if (c === undefined || c.name === undefined) {
				return "";
			}

			return `${c.name} (${c.genre}) ${c.lc} - ${c.ct}`;
		},
	},

	methods: {
		...mapActions(["setPlayerStatus"]),

		tooglePlay() {
			var player = document.getElementById("player");

			if (this.getPlayerStatus == "paused") {
				//play
				this.playerIcon = "pause";

				player.pause();
				player.src = this.getPlayerSource;
				player.load();
				player.play();

				this.setPlayerStatus("playing");
			} else {
				//pause
				this.playerIcon = "play_arrow";
				player.pause();
				this.setPlayerStatus("paused");
			}
		},
	},

	watch: {
		getPlayerSource: val => {
			try {
				var player = document.getElementById("player");

				var url = new URL(val);

				player.pause();

				player.src = url.href;

				player.load();

				player.play().then(
					() => {},
					() => {
						player.src = url.href + ";";
						player.load();
						player.play();
					}
				);
			} catch (e) {
				alert("cannot play..." + e);
			}
		},
	},
};
</script>

<style lang="scss"></style>
