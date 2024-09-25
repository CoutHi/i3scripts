# i3scripts

## Example i3blocks config

```
[playback]
command=echo "Now Playing:" $([path_to_git_repo]/i3scripts/now_playing)
interval=1

[i3volume]
label=  
command=[path_to_git_repo]/i3scripts/get_volume.sh
interval=once
signal=10
```

To make i3volume change with the volume, you'd append the line:
```
&& pkill -SIGRTMIN+10 i3blocks
```

To where you define the shortcuts to handle volume.

### Example

```bindsym XF86AudioLowerVolume exec --no-startup-id pactl set-sink-volume @DEFAULT_SINK@ -5% && $refresh_i3status && pkill -SIGRTMIN+10 i3blocks```

Above will send the signal 10 to i3blocks whenever you lower the volume, which will refresh the output on your bar.

If you add the lines in volume_control_buttons.txt to your i3blocks config, you'll have two buttons to increase and decrease the audio, which also send the needed signal to the volume level block that updates it.
