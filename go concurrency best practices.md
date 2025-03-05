### **Best practices per la concorrenza in Go**  

1. **Evitare goroutine nelle librerie**  
   - Se stai sviluppando una libreria, lascia che sia il consumatore a gestire la concorrenza.  
   - L’uso forzato di goroutine nella libreria può rendere più difficile la sincronizzazione dei dati per chi la utilizza.  
   - Un'eccezione è quando una funzione restituisce un **channel** per i risultati: in tal caso, la concorrenza è nascosta all’utente finale.  

2. **Gestire correttamente la terminazione delle goroutine**  
   - Non lasciare goroutine attive senza un meccanismo di arresto.  
   - Una goroutine che continua a funzionare indefinitamente può consumare risorse inutilmente e, nel tempo, portare a crash dell’applicazione.  
   - L’uso di **canali** può aiutare a segnalare quando una goroutine deve terminare.  

3. **Verificare le race condition a tempo di compilazione**  
   - Utilizza strumenti come **`go run -race`** per rilevare problemi di concorrenza prima dell’esecuzione.  
   - Le race condition possono essere difficili da individuare a runtime, quindi testarle durante lo sviluppo è fondamentale.  

Queste pratiche aiutano a scrivere codice concorrente più sicuro, efficiente e manutenibile in Go.