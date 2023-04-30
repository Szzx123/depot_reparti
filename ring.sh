# Fonction de nettoyage
nettoyer () {
  # Suppression des processus de l'application app
  killall app 2> /dev/null

  # Suppression des processus de l'application ctl
  killall ctl 2> /dev/null

  # Suppression des processus tee et cat
  killall tee 2> /dev/null
  killall cat 2> /dev/null

  # Suppression des tubes nommés
  \rm -f /tmp/in* /tmp/out*
  \rm -f error.log
  exit 0
}

# Appel de la fonction nettoyer à la réception d'un signal


mkfifo /tmp/in_A1 /tmp/out_A1
mkfifo /tmp/in_C1 /tmp/out_C1

mkfifo /tmp/in_A2 /tmp/out_A2
mkfifo /tmp/in_C2 /tmp/out_C2

mkfifo /tmp/in_A3 /tmp/out_A3
mkfifo /tmp/in_C3 /tmp/out_C3

./app -n A1 -p 8080 < /tmp/in_A1 > /tmp/out_A1 &
./ctl -n C1 < /tmp/in_C1 > /tmp/out_C1 &

./app -n A2 -p 8081 < /tmp/in_A2 > /tmp/out_A2 &
./ctl -n C2 < /tmp/in_C2 > /tmp/out_C2 &

./app -n A3 -p 8082 < /tmp/in_A3 > /tmp/out_A3 &
./ctl -n C3 < /tmp/in_C3 > /tmp/out_C3 &

cat /tmp/out_A1 > /tmp/in_C1 &
cat /tmp/out_C1 | tee /tmp/in_A1 > /tmp/in_C2 &

cat /tmp/out_A2 > /tmp/in_C2 &
cat /tmp/out_C2 | tee /tmp/in_A2 > /tmp/in_C3 &

cat /tmp/out_A3 > /tmp/in_C3 &
cat /tmp/out_C3 | tee /tmp/in_A3 > /tmp/in_C1 &



echo "+ INT QUIT TERM => nettoyer"
trap nettoyer INT QUIT TERM
echo "+ Attente de Crtl C pendant 1h..."
sleep 3600
nettoyer