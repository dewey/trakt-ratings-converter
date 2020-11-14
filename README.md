# Trakt to IMDb Rating converter

This tiny tool fetches your Trakt.tv movie ratings and exports a CSV file with the same format as the one you can export from IMDB itself. This can be useful if you want to import it on that one popcorn-themed movie site.

The file from IMDb can usually be exported here: `https://www.imdb.com/user/ur{your user ID}/ratings/export`

## Usage

Open `run_develop.sh` and add your Trakt.tv username and Client ID (You can find that one in [your profile](https://trakt.tv/oauth/applications)).

Then execute `./run_develop.sh`. This will result in a `trakt-ratings.csv` file being generated in the same directory.