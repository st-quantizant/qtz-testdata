      *----------------------------------------------------------------*
      * Sample COBOL program demonstrating weak cryptographic usage    *
      * detected by SAST regex and LLM-based analysis.                 *
      *----------------------------------------------------------------*
       IDENTIFICATION DIVISION.
       PROGRAM-ID. CRYPTO-SAMPLE.

       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 WS-HASH-ALG          PIC X(10).
       01 WS-DATA              PIC X(1024).
       01 WS-DIGEST            PIC X(64).
       01 WS-RETURN-CODE       PIC S9(9) COMP.

       PROCEDURE DIVISION.

      *-- Weak: MD5 -- SAST should flag this
       WEAK-HASH-MD5.
           MOVE "MD5" TO WS-HASH-ALG
           CALL 'MD5-INIT'
               USING WS-DATA, WS-DIGEST, WS-RETURN-CODE
           STOP RUN.

      *-- Safe: SHA-256
       STRONG-HASH-SHA256.
           MOVE "SHA-256" TO WS-HASH-ALG
           CALL 'SHA256-FINAL'
               USING WS-DATA, WS-DIGEST, WS-RETURN-CODE
           STOP RUN.
