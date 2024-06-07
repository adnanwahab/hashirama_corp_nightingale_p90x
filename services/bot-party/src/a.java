       ExoPlayer player = new ExoPlayer.Builder(context).build();
// Set the media item to be played.
        player.setMediaItem(MediaItem.fromUri(rtspUri));
// Prepare the player.
        player.prepare();

        MediaSource mediaSource =
                new RtspMediaSource.Factory().createMediaSource(MediaItem.fromUri(rtspUri));
// Create a player instance.
        ExoPlayer player = new ExoPlayer.Builder(context).build();
// Set the media source to be played.
        player.setMediaSource(mediaSource);
// Prepare the player.
        player.prepare();

        MediaSource mediaSource =
                new RtspMediaSource.Factory()
                        .setSocketFactory(...)
        .createMediaSource(MediaItem.fromUri(http://192.168.1.202:5432));
