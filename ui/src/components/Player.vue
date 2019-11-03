<template>
	<div>
		<md-button
			@click="tooglePlay"
			class="md-icon-button md-theme-footer md-primary"
		>
			<md-icon>{{ playerIcon }}</md-icon>
		</md-button>
		<md-button class="md-icon-button md-theme-footer md-primary">
			<md-icon>volume_down</md-icon>
		</md-button>
		<md-button class="md-icon-button md-theme-footer md-primary">
			<md-icon>volume_up</md-icon>
		</md-button>
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
			playerIcon: "play_arrow"
		};
	},

	props: {
		source: {
			type: String,
			default: "http://dream01.gr:8000/stream"
		}
	},

	created() {},

	computed: {
		...mapGetters(["getPlayerStatus", "getPlayerSource"])
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
		}
	},

	watch: {
		getPlayerSource: val => {
			var player = document.getElementById("player");

			var url = new URL(val);

			player.pause();
			player.src = url.href;
			player.load();
			player.play().then(
				res => {
					console.log(res);
				},
				err => {
					console.log(err);

					player.src = url.href + ";";
					player.load();
					player.play();
				}
			);
		}
	}
};
</script>

<style lang="scss">
</style>